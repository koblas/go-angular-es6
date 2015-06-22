function MainRegister(app) {
    app.config(($stateProvider) => {
        $stateProvider
            .state('app.index_', {
                url: '',
                views: {
                    'content@app': {
                        template: require('../../partials/index.html'),
                        controller: "IndexController",
                        authenticate: false
                    }
                }
            });
        $stateProvider
            .state('app.index', {
                url: '/',
                views: {
                    'content@app': {
                        template: require('../../partials/index.html'),
                        controller: "IndexController",
                        authenticate: false
                    }
                }
            });
    })

    app.controller('HeaderController', HeaderController)
    app.controller('IndexController', IndexController)
}

class IndexController {
    /*@ngInject*/
    constructor(AuthService, $state) {
        if (AuthService.isAuthenticated())
            $state.go('app.dashboard')
    }
}

class HeaderController {
    /*@ngInject*/
    constructor(AuthService, $location) {
        this.auth = AuthService;
        this.$location = $location
    }

    logout() {
        this.auth.logout();
        this.$location.path('/');
    }
}

export { MainRegister }
