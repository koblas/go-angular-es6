describe("App Module:", function() {
    var module;

    beforeEach(function() {
        module = angular.module('app');
    });

    it("should be registered", function() {
        expect(module).not.toEqual(null);
    });

    /*
    var RegisterController, scope;

    beforeEach(inject(function($rootScope, $controller) {
        scope = $rootScope.$new();
        RegisterController = $controller('RegisterController', { $scope: scope });
    }));

    //
    //  Start Tests
    //
    it('say hello', function() {
        expect(scope.username).toEqual('')
    })
    */

});
