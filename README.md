# CLI Todo List

A simple yet powerful command-line Todo List application written in Go. This application allows users to manage their tasks with user authentication, priority levels, and data persistence.

## Features

- **User Authentication**: Create accounts and log in securely
- **Todo Management**: Add, edit, delete, and mark todos as complete
- **Priority Levels**: Assign priority levels (1-4) to your tasks
- **Visual Indicators**: Color-coded priorities and completion status
- **Data Persistence**: Todos and user data are saved between sessions
- **User-specific Todos**: Each user can only view and manage their own todos

## Installation

### Prerequisites

- Go 1.18 or higher

### Steps

1. Clone the repository:
```bash
git clone https://github.com/Blaze5333/todo_cli_app.git
cd cli_todo_list
```

2. Build the application:
```bash
go build
```

3. Run the application:
```bash
./cli_todo_list
```

## Usage

### Authentication

When you start the application, you'll be presented with the following options:
1. Login
2. Create a new user
3. Exit

If you're a new user, select option 2 to create an account. If you already have an account, select option 1 to log in.

### Managing Todos

After logging in, you can:

1. **Add a new todo**: Create a new task with a title and priority level (1-4)
2. **Edit an existing todo**: Change the title of an existing task
3. **Delete a todo**: Remove a task from your list
4. **Complete a todo**: Mark a task as completed
5. **Display all todos**: View all your tasks with their details
6. **Exit**: Exit the application

### Priority Levels

Todos are color-coded according to their priority:
- 🟢 (1): Low priority
- 🔵 (2): Medium priority
- 🟡 (3): High priority
- 🔴 (4): Urgent

### Task Status

- ❌: Task not completed
- ✅: Task completed

## Data Storage

The application stores data in two JSON files:
- `users.json`: Contains user credentials
- `todos.json`: Contains all todos organized by user

## Dependencies

This project uses the following external packages:
- [github.com/aquasecurity/table](https://github.com/aquasecurity/table): For formatted table output

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Future Improvements

- Command-line flag support for quick actions
- Due dates for todos
- Categories/tags for todos
- Search and filter functionality
- Export/import todos
