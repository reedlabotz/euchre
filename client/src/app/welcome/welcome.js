angular.module('welcome', [])
    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/welcome', {
            templateUrl: 'welcome/welcome.tpl.html',
            controller: 'WelcomeCtrl'
        });
    }])
    .controller('WelcomeCtrl', ['$scope', '$location', function($scope, $location) {
        $scope.startGame = function() {
            $location.path("/game/abcd");
        };
    }]);
