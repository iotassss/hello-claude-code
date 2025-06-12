class TodoApp {
    constructor() {
        this.tasks = [];
        this.categories = [];
        this.currentEditId = null;
        this.currentDeleteId = null;
        this.filters = {
            status: '',
            priority: '',
            category: '',
            search: ''
        };
        this.sort = {
            by: 'created_at',
            order: 'desc'
        };
        
        this.init();
    }

    async init() {
        this.bindEvents();
        await this.loadCategories();
        await this.loadTasks();
        await this.loadStats();
        this.setupTheme();
    }

    bindEvents() {
        // Modal events
        document.getElementById('add-task-btn').addEventListener('click', () => this.showTaskModal());
        document.getElementById('close-modal').addEventListener('click', () => this.hideTaskModal());
        document.getElementById('cancel-btn').addEventListener('click', () => this.hideTaskModal());
        document.getElementById('task-form').addEventListener('submit', (e) => this.handleTaskSubmit(e));
        
        // Delete modal events
        document.getElementById('close-delete-modal').addEventListener('click', () => this.hideDeleteModal());
        document.getElementById('cancel-delete-btn').addEventListener('click', () => this.hideDeleteModal());
        document.getElementById('confirm-delete-btn').addEventListener('click', () => this.confirmDelete());
        
        // Filter events
        document.querySelectorAll('input[name="status"]').forEach(radio => {
            radio.addEventListener('change', (e) => {
                this.filters.status = e.target.value;
                this.loadTasks();
            });
        });
        
        document.querySelectorAll('input[name="priority"]').forEach(radio => {
            radio.addEventListener('change', (e) => {
                this.filters.priority = e.target.value;
                this.loadTasks();
            });
        });
        
        document.getElementById('category-filter').addEventListener('change', (e) => {
            this.filters.category = e.target.value;
            this.loadTasks();
        });
        
        // Search events
        let searchTimeout;
        document.getElementById('search-input').addEventListener('input', (e) => {
            clearTimeout(searchTimeout);
            searchTimeout = setTimeout(() => {
                this.filters.search = e.target.value;
                this.loadTasks();
            }, 300);
        });
        
        // Sort events
        document.getElementById('sort-by').addEventListener('change', (e) => {
            this.sort.by = e.target.value;
            this.loadTasks();
        });
        
        document.getElementById('sort-order').addEventListener('change', (e) => {
            this.sort.order = e.target.value;
            this.loadTasks();
        });
        
        // Theme toggle
        document.getElementById('toggle-theme').addEventListener('click', () => this.toggleTheme());
        
        // Close modals on outside click
        document.getElementById('task-modal').addEventListener('click', (e) => {
            if (e.target.id === 'task-modal') this.hideTaskModal();
        });
        
        document.getElementById('delete-modal').addEventListener('click', (e) => {
            if (e.target.id === 'delete-modal') this.hideDeleteModal();
        });
    }

    async apiCall(url, method = 'GET', data = null) {
        const options = {
            method,
            headers: {
                'Content-Type': 'application/json'
            }
        };
        
        if (data) {
            options.body = JSON.stringify(data);
        }
        
        try {
            const response = await fetch(url, options);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return await response.json();
        } catch (error) {
            console.error('API call failed:', error);
            this.showNotification('An error occurred. Please try again.', 'error');
            throw error;
        }
    }

    async loadCategories() {
        try {
            const response = await this.apiCall('/api/v1/categories');
            this.categories = response.data;
            this.populateCategorySelects();
        } catch (error) {
            console.error('Failed to load categories:', error);
        }
    }

    populateCategorySelects() {
        const categoryFilter = document.getElementById('category-filter');
        const taskCategory = document.getElementById('task-category');
        
        // Clear existing options (except "All Categories" and "No Category")
        categoryFilter.innerHTML = '<option value="">All Categories</option>';
        taskCategory.innerHTML = '<option value="">No Category</option>';
        
        this.categories.forEach(category => {
            const filterOption = document.createElement('option');
            filterOption.value = category.id;
            filterOption.textContent = category.name;
            categoryFilter.appendChild(filterOption);
            
            const taskOption = document.createElement('option');
            taskOption.value = category.id;
            taskOption.textContent = category.name;
            taskCategory.appendChild(taskOption);
        });
    }

    async loadTasks() {
        const loading = document.getElementById('loading');
        loading.style.display = 'block';
        
        try {
            const params = new URLSearchParams();
            
            if (this.filters.status) params.append('status', this.filters.status);
            if (this.filters.priority) params.append('priority', this.filters.priority);
            if (this.filters.category) params.append('category_id', this.filters.category);
            if (this.filters.search) params.append('search', this.filters.search);
            params.append('sort', this.sort.by);
            params.append('order', this.sort.order);
            
            const response = await this.apiCall(`/api/v1/tasks?${params}`);
            this.tasks = response.data;
            this.renderTasks();
            await this.loadStats();
        } catch (error) {
            console.error('Failed to load tasks:', error);
        } finally {
            loading.style.display = 'none';
        }
    }

    async loadStats() {
        try {
            const response = await this.apiCall('/api/v1/tasks/stats');
            const stats = response.data;
            
            document.getElementById('total-tasks').textContent = stats.total;
            document.getElementById('pending-tasks').textContent = stats.pending + stats.in_progress;
            document.getElementById('completed-tasks').textContent = stats.completed;
        } catch (error) {
            console.error('Failed to load stats:', error);
        }
    }

    renderTasks() {
        const container = document.getElementById('tasks-container');
        const loading = document.getElementById('loading');
        
        if (this.tasks.length === 0) {
            container.innerHTML = `
                <div class="empty-state">
                    <h3>No tasks found</h3>
                    <p>Create your first task to get started!</p>
                    <button class="btn btn-primary" onclick="app.showTaskModal()">Add Task</button>
                </div>
            `;
            return;
        }
        
        container.innerHTML = this.tasks.map(task => this.renderTaskCard(task)).join('');
    }

    renderTaskCard(task) {
        const dueDate = task.due_date ? new Date(task.due_date).toLocaleDateString() : null;
        const isOverdue = task.due_date && new Date(task.due_date) < new Date() && task.status !== 'completed';
        const category = task.category ? task.category : null;
        
        return `
            <div class="task-card priority-${task.priority} ${task.status === 'completed' ? 'completed' : ''}" data-id="${task.id}">
                <div class="task-header">
                    <input type="checkbox" class="task-checkbox" ${task.status === 'completed' ? 'checked' : ''} 
                           onchange="app.toggleTaskStatus(${task.id})">
                    <div class="task-content">
                        <div class="task-title">${this.escapeHtml(task.title)}</div>
                        ${task.description ? `<div class="task-description">${this.escapeHtml(task.description)}</div>` : ''}
                        <div class="task-meta">
                            <span class="priority-badge ${task.priority}">${task.priority}</span>
                            <span class="status-badge ${task.status}">${task.status.replace('_', ' ')}</span>
                            ${category ? `<span class="task-category" style="background-color: ${category.color}">${category.name}</span>` : ''}
                            ${dueDate ? `<span class="${isOverdue ? 'text-danger' : ''}">📅 ${dueDate}</span>` : ''}
                        </div>
                    </div>
                    <div class="task-actions">
                        <button class="btn btn-small btn-secondary" onclick="app.editTask(${task.id})">Edit</button>
                        <button class="btn btn-small btn-danger" onclick="app.showDeleteModal(${task.id})">Delete</button>
                    </div>
                </div>
            </div>
        `;
    }

    showTaskModal(taskId = null) {
        this.currentEditId = taskId;
        const modal = document.getElementById('task-modal');
        const form = document.getElementById('task-form');
        const title = document.getElementById('modal-title');
        
        if (taskId) {
            title.textContent = 'Edit Task';
            const task = this.tasks.find(t => t.id === taskId);
            if (task) {
                document.getElementById('task-title').value = task.title;
                document.getElementById('task-description').value = task.description || '';
                document.getElementById('task-priority').value = task.priority;
                document.getElementById('task-status').value = task.status;
                document.getElementById('task-category').value = task.category_id || '';
                
                if (task.due_date) {
                    const date = new Date(task.due_date);
                    const formattedDate = date.toISOString().slice(0, 16);
                    document.getElementById('task-due-date').value = formattedDate;
                } else {
                    document.getElementById('task-due-date').value = '';
                }
            }
        } else {
            title.textContent = 'Add New Task';
            form.reset();
            document.getElementById('task-priority').value = 'medium';
            document.getElementById('task-status').value = 'pending';
        }
        
        modal.classList.add('show');
        document.getElementById('task-title').focus();
    }

    hideTaskModal() {
        const modal = document.getElementById('task-modal');
        modal.classList.remove('show');
        this.currentEditId = null;
    }

    async handleTaskSubmit(e) {
        e.preventDefault();
        
        const formData = {
            title: document.getElementById('task-title').value.trim(),
            description: document.getElementById('task-description').value.trim(),
            priority: document.getElementById('task-priority').value,
            due_date: document.getElementById('task-due-date').value || null,
            category_id: document.getElementById('task-category').value || null
        };
        
        if (this.currentEditId) {
            formData.status = document.getElementById('task-status').value;
        }
        
        if (!formData.title) {
            this.showNotification('Task title is required', 'error');
            return;
        }
        
        try {
            if (this.currentEditId) {
                await this.apiCall(`/api/v1/tasks/${this.currentEditId}`, 'PUT', formData);
                this.showNotification('Task updated successfully', 'success');
            } else {
                await this.apiCall('/api/v1/tasks', 'POST', formData);
                this.showNotification('Task created successfully', 'success');
            }
            
            this.hideTaskModal();
            await this.loadTasks();
        } catch (error) {
            console.error('Failed to save task:', error);
        }
    }

    editTask(taskId) {
        this.showTaskModal(taskId);
    }

    showDeleteModal(taskId) {
        this.currentDeleteId = taskId;
        const modal = document.getElementById('delete-modal');
        modal.classList.add('show');
    }

    hideDeleteModal() {
        const modal = document.getElementById('delete-modal');
        modal.classList.remove('show');
        this.currentDeleteId = null;
    }

    async confirmDelete() {
        if (!this.currentDeleteId) return;
        
        try {
            await this.apiCall(`/api/v1/tasks/${this.currentDeleteId}`, 'DELETE');
            this.showNotification('Task deleted successfully', 'success');
            this.hideDeleteModal();
            await this.loadTasks();
        } catch (error) {
            console.error('Failed to delete task:', error);
        }
    }

    async toggleTaskStatus(taskId) {
        try {
            await this.apiCall(`/api/v1/tasks/${taskId}/toggle`, 'PATCH');
            await this.loadTasks();
        } catch (error) {
            console.error('Failed to toggle task status:', error);
            // Revert checkbox state
            const checkbox = document.querySelector(`[data-id="${taskId}"] .task-checkbox`);
            if (checkbox) {
                checkbox.checked = !checkbox.checked;
            }
        }
    }

    setupTheme() {
        const savedTheme = localStorage.getItem('theme') || 'light';
        document.documentElement.setAttribute('data-theme', savedTheme);
        this.updateThemeButton(savedTheme);
    }

    toggleTheme() {
        const currentTheme = document.documentElement.getAttribute('data-theme');
        const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
        
        document.documentElement.setAttribute('data-theme', newTheme);
        localStorage.setItem('theme', newTheme);
        this.updateThemeButton(newTheme);
    }

    updateThemeButton(theme) {
        const button = document.getElementById('toggle-theme');
        button.textContent = theme === 'dark' ? '☀️' : '🌙';
    }

    showNotification(message, type = 'info') {
        // Create notification element
        const notification = document.createElement('div');
        notification.className = `notification notification-${type}`;
        notification.innerHTML = `
            <span>${message}</span>
            <button onclick="this.parentElement.remove()">&times;</button>
        `;
        
        // Add to page
        document.body.appendChild(notification);
        
        // Auto remove after 5 seconds
        setTimeout(() => {
            if (notification.parentElement) {
                notification.remove();
            }
        }, 5000);
    }

    escapeHtml(text) {
        const div = document.createElement('div');
        div.textContent = text;
        return div.innerHTML;
    }
}

// Initialize app when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    window.app = new TodoApp();
});

// Add notification styles
const notificationStyles = `
.notification {
    position: fixed;
    top: 20px;
    right: 20px;
    padding: 1rem 1.5rem;
    border-radius: 6px;
    color: white;
    font-weight: 500;
    z-index: 2000;
    display: flex;
    align-items: center;
    gap: 1rem;
    max-width: 400px;
    animation: slideInRight 0.3s ease;
}

.notification-success {
    background: #2ecc71;
}

.notification-error {
    background: #e74c3c;
}

.notification-info {
    background: #3498db;
}

.notification button {
    background: none;
    border: none;
    color: white;
    font-size: 1.2rem;
    cursor: pointer;
    padding: 0;
    width: 24px;
    height: 24px;
    display: flex;
    align-items: center;
    justify-content: center;
}

@keyframes slideInRight {
    from {
        transform: translateX(100%);
        opacity: 0;
    }
    to {
        transform: translateX(0);
        opacity: 1;
    }
}
`;

// Add styles to document
const styleSheet = document.createElement('style');
styleSheet.textContent = notificationStyles;
document.head.appendChild(styleSheet);