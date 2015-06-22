function TodoRegister(app) {
    app.config(($stateProvider) => {
        $stateProvider
            .state('app.dashboard', {
                url: '/dash',
                authenticate: true,
                views: {
                    'content@app': {
                        template: require('../../partials/dashboard.html'),
                        // controller: "DashboardController",
                    }
                }
            })
            .state('app.todo', {
                url: '/todo',
                authenticate: true,
                views: {
                    'content@app': {
                        template: require('../../partials/todo.html'),
                    }
                }
            });
    })

    app.controller('TodoController', TodoController)
}

class TodoController {
    /*@ngInject*/
    constructor(Restangular) {
        this.todos = [];
        this.loaded = false;

        this.Todo = Restangular.all('todo')
        this.newTodoTitle = ''

        this.Todo.getList().then(
            (todos) => {
                this.loaded = true
                this.todos = todos
            }
        )
    }

    addTodo(title) {
        if (title) {
            this.Todo.post({'title': title }).then(
                (todo) => {
                    this.newTodoTitle = '';
                    if (todo)
                        this.todos.push(todo);
                }, (err) => {
                    return alert(err.data.message || "an error occurred");
                }
            )
        }
    }

    changeCompleted(todo) {
        todo.put().then(null, (err) => {
            return alert(err.data.message || (err.errors && err.errors.completed) || "an error occurred");
        })
    }

    removeCompletedItems() {
        this.todos.forEach((todo) => {
            if (!todo.completed)
                return;

            todo.remove().then(() => {
                this.todos = _.without(this.todos, todo);
            }, (err) => {
                return alert(err.data.message || (err.errors && err.errors.completed) || "an error occurred");
            })
        })
    }
}

export { TodoRegister }
