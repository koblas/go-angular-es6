function AuthRegister(app) {
    app.config(($stateProvider) => {
        //  States
        $stateProvider
            .state('app.auth', {
                abstract: true,
                url: "/auth",
                template: '<ui-view/>'
            })
            .state('app.auth.register', {
                url: "/register?next",
                views: {
                    'content@app': {
                        template: require("../../partials/auth/register.html"),
                        authenticate: false
                    }
                }
            })
            .state('app.auth.login', {
                url: "/login?next",
                views: {
                    'content@app': {
                        template: require("../../partials/auth/login.html"),
                        authenticate: false
                    }
                }
            })
            .state('app.auth.logout', {
                url: "/logout",
                // template: require("../../partials/auth/logout.html")
                views: {
                    'content@app': {
                        template: "<div></div>",
                        controller: "LogoutController",
                        authenticate: false
                    }
                }
            })
    })

    app.controller('LoginController', LoginController)
    app.controller('RegisterController', RegisterController)
}

class LoginController {
    constructor($location, AuthService, $stateParams) {
        this.AuthService = AuthService
        this.$location = $location

        this.email = ""
        this.error = ""
        this.next = $stateParams.next ? $stateParams.next : "/"
        console.log($stateParams)
    }

    login() {
        if (!this.email) {
            this.error = "Invalid Email Address"
            return
        }

        if (!this.password) {
            this.error = "Invalid Password"
            return
        }

        console.log("NEXT = ", this.next)

        this.AuthService.login(this.email, this.password)
            .then(() => this.$location.path(this.next))
            .catch((err) => console.log("Auth Error", err))
    }
}

LoginController.$inject = ['$location', 'AuthService', '$stateParams'];


class RegisterController {
    constructor($location, AuthService, $stateParams) {
        this.AuthService = AuthService
        this.$location = $location

        this.password = ""
        this.username = ""
        this.email = ""
        this.error = ""
        this.next = $stateParams.next ? $stateParams.next : "/"
    }

    register() {
        if (!this.email) {
            this.error = "Invalid Email Address"
            return
        }
        if (!this.username) {
            this.error = "Invalid Username"
            return
        }
        if (!this.password) {
            this.error = "Invalid Password"
            return
        }

        this.AuthService.register(this.email, this.password, { username: this.username })
            .then(() => this.$location.path(this.next))
            .catch((err) => this.error = err)
    }
}

RegisterController.$inject = ['$location', 'AuthService', '$stateParams'];

export { AuthRegister }
