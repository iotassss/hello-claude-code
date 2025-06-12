# Screen Design Specification

## Overview

UI/UX design specification for the TODO App web interface with responsive design for desktop, tablet, and mobile devices.

## Design Principles

- **Minimalist**: Clean, uncluttered interface
- **Responsive**: Works on all device sizes
- **Accessible**: WCAG 2.1 AA compliance
- **Intuitive**: Clear navigation and actions
- **Consistent**: Unified color scheme and typography

## Color Scheme

### Primary Colors
- **Primary Blue**: #3498db (buttons, links, highlights)
- **Success Green**: #2ecc71 (completed tasks, success messages)
- **Warning Orange**: #f39c12 (medium priority, warnings)
- **Danger Red**: #e74c3c (high priority, delete actions)
- **Dark Gray**: #2c3e50 (text, headers)
- **Light Gray**: #ecf0f1 (backgrounds, borders)
- **White**: #ffffff (cards, modals)

### Priority Colors
- **Low**: #95a5a6 (gray)
- **Medium**: #f39c12 (orange)
- **High**: #e67e22 (dark orange)
- **Critical**: #e74c3c (red)

## Typography

### Font Family
- **Primary**: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif
- **Monospace**: 'Fira Code', 'Courier New', monospace (for code/timestamps)

### Font Sizes
- **H1**: 2rem (32px) - Page titles
- **H2**: 1.5rem (24px) - Section headers
- **H3**: 1.25rem (20px) - Card titles
- **Body**: 1rem (16px) - Regular text
- **Small**: 0.875rem (14px) - Meta information
- **Tiny**: 0.75rem (12px) - Labels, timestamps

## Layout Structure

### Desktop Layout (1024px+)
```
┌─────────────────────────────────────────────────────────┐
│                    Header (60px)                        │
├─────────────────────────────────────────────────────────┤
│ Sidebar │                Main Content                   │
│ (280px) │                 (flexible)                    │
│         │                                               │
│         │                                               │
│         │                                               │
└─────────┴───────────────────────────────────────────────┘
```

### Mobile Layout (< 768px)
```
┌─────────────────────────────────────┐
│          Header (60px)             │
├─────────────────────────────────────┤
│                                     │
│         Main Content                │
│         (full width)                │
│                                     │
│                                     │
├─────────────────────────────────────┤
│      Bottom Navigation (60px)       │
└─────────────────────────────────────┘
```

## Screen Specifications

### 1. Main Dashboard

#### Desktop View
```
┌─────────────────────────────────────────────────────────┐
│ [Logo] TODO App           [Search] [Profile] [Settings] │
├─────────────────────────────────────────────────────────┤
│ Filters  │ ┌─ Quick Stats ─────────────────────────┐   │
│ □ All    │ │ Total: 25 | Pending: 15 | Done: 10    │   │
│ □ Pending│ └─────────────────────────────────────────┘   │
│ □ Done   │                                             │
│ ────────│ ┌─ Add New Task ──────────────────────────┐   │
│ Priority │ │ [Title input field            ] [Add]  │   │
│ □ Critical│ └─────────────────────────────────────────┘   │
│ □ High   │                                             │
│ □ Medium │ ┌─ Task List ─────────────────────────────┐   │
│ □ Low    │ │ ☐ Buy groceries      [Edit] [Delete]   │   │
│ ──────── │ │   Due: Today • Work • High             │   │
│ Category │ │ ☐ Finish project     [Edit] [Delete]   │   │
│ • Work   │ │   Due: Tomorrow • Work • Critical      │   │
│ • Personal│ │ ☑ Call dentist       [Edit] [Delete]   │   │
│ • Health │ │   Completed • Personal • Low           │   │
│          │ └─────────────────────────────────────────┘   │
└─────────┴─────────────────────────────────────────────┘
```

#### Key Elements:
- **Header**: Logo, search bar, user menu
- **Sidebar**: Filters, categories, quick actions
- **Main Area**: Stats, quick add, task list
- **Task Cards**: Checkbox, title, metadata, actions

### 2. Task Detail Modal

```
┌─────────────────────────────────────────┐
│ Edit Task                          [×]  │
├─────────────────────────────────────────┤
│ Title: [Buy groceries for dinner      ] │
│                                         │
│ Description:                            │
│ [Need to get vegetables, fruits, and   ]│
│ [bread for tonight's dinner           ]│
│                                         │
│ Due Date: [2024-01-15] [14:30]         │
│                                         │
│ Priority: ○ Low ○ Medium ● High ○ Critical│
│                                         │
│ Category: [Work ▼]                      │
│                                         │
│ Status: ○ Pending ● In Progress ○ Done  │
│                                         │
│         [Cancel]  [Save Changes]        │
└─────────────────────────────────────────┘
```

### 3. Add Task Modal

```
┌─────────────────────────────────────────┐
│ Add New Task                       [×]  │
├─────────────────────────────────────────┤
│ Title: [                              ] │
│                                         │
│ Description:                            │
│ [                                     ] │
│ [                                     ] │
│                                         │
│ Due Date: [        ] [     ] (Optional) │
│                                         │
│ Priority: ● Low ○ Medium ○ High ○ Critical│
│                                         │
│ Category: [Personal ▼]                  │
│                                         │
│         [Cancel]      [Add Task]        │
└─────────────────────────────────────────┘
```

### 4. Mobile Dashboard

```
┌─────────────────────────────────────┐
│ ☰ TODO App        🔍     👤        │
├─────────────────────────────────────┤
│ [+ Add Task                       ] │
│                                     │
│ ┌─ Today's Tasks ──────────────────┐ │
│ │ ☐ Buy groceries                 │ │
│ │   Work • High • Due today       │ │
│ │                                 │ │
│ │ ☐ Team meeting                  │ │
│ │   Work • Medium • 2:00 PM       │ │
│ └─────────────────────────────────────┘ │
│                                     │
│ ┌─ Upcoming ──────────────────────┐ │
│ │ ☐ Dentist appointment           │ │
│ │   Personal • High • Tomorrow    │ │
│ └─────────────────────────────────────┘ │
│                                     │
├─────────────────────────────────────┤
│ [All] [Pending] [Done] [Categories] │
└─────────────────────────────────────┘
```

## Component Specifications

### Task Card Component

#### States:
- **Default**: Regular task display
- **Completed**: Strikethrough text, muted colors
- **Overdue**: Red accent, warning icon
- **High Priority**: Red left border
- **Selected**: Blue background highlight

#### Hover Effects:
- Subtle shadow elevation
- Action buttons fade in
- Background color change

### Button Specifications

#### Primary Button
- Background: #3498db
- Text: White
- Border radius: 6px
- Padding: 12px 24px
- Hover: #2980b9

#### Secondary Button
- Background: Transparent
- Text: #3498db
- Border: 1px solid #3498db
- Hover: Background #3498db, Text white

#### Danger Button
- Background: #e74c3c
- Text: White
- Hover: #c0392b

### Form Elements

#### Input Fields
- Border: 1px solid #ddd
- Border radius: 4px
- Padding: 12px
- Focus: Blue border (#3498db)
- Font size: 16px (prevents zoom on mobile)

#### Select Dropdowns
- Custom styled with arrow icon
- Same styling as input fields
- Dropdown shadow: 0 4px 12px rgba(0,0,0,0.15)

#### Checkboxes
- Custom styled with brand colors
- Checked: #3498db background
- Size: 20px × 20px

## Responsive Breakpoints

### Desktop Large (1200px+)
- Sidebar: 280px fixed
- Task cards: 3 columns in grid view
- Full feature set visible

### Desktop (1024px - 1199px)
- Sidebar: 240px fixed
- Task cards: 2 columns in grid view
- Condensed navigation

### Tablet (768px - 1023px)
- Sidebar: Collapsible overlay
- Task cards: 2 columns, then 1 column
- Touch-friendly sizing

### Mobile (< 768px)
- Bottom navigation
- Full-width task cards
- Simplified interface
- Swipe gestures for actions

## Animation Specifications

### Transitions
- Default: 0.3s ease-in-out
- Hover effects: 0.2s ease
- Modal open/close: 0.25s ease-out
- Task completion: 0.4s ease (with strikethrough animation)

### Loading States
- Skeleton screens for task list
- Spinner for form submissions
- Progressive loading for images

## Accessibility Features

### Keyboard Navigation
- Tab order: logical flow
- Enter/Space: activate buttons
- Escape: close modals
- Arrow keys: navigate lists

### Screen Reader Support
- ARIA labels on all interactive elements
- Proper heading hierarchy
- Focus indicators
- Status announcements

### Color Contrast
- Minimum 4.5:1 for normal text
- Minimum 3:1 for large text
- High contrast mode support

## Dark Mode Support

### Color Adjustments
- Background: #1a1a1a
- Cards: #2d2d2d
- Text: #ffffff
- Borders: #404040
- Maintains color contrast ratios

### Implementation
- CSS custom properties
- System preference detection
- Toggle switch in settings
- Persistent user preference