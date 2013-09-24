var gameModule = angular.module('game', ['common']);
gameModule.config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/game/:id', {
        templateUrl:'game/game.tpl.html',
        controller:'GameCtrl'
    });
}]);
gameModule.controller('GameCtrl', ['$scope', function($scope) {
    $scope.game = {
        Teams: [
            {
                Players: [
                    {
                        Name: 'Reed',
                        Id: 'abcd-abcd-abcd-abcd',
                        Hand: [0,1,2,3,4]
                    },
                    {
                        Name: 'Adina',
                        Id: 'abcd-abcd-abcd-abcd',
                        Hand: [-1, -1, -1, -1, -1]
                    }
                ],
                Score: 4,
                HandsWon: 2
            },
            {
                Players: [
                    {
                        Name: 'Joe',
                        Id: 'abcd-abcd-abcd-abcd',
                        Hand: [-1, -1, -1, -1, -1]
                    },
                    {
                        Name: 'Serena',
                        Id: 'abcd-abcd-abcd-abcd',
                        Hand: [-1, -1, -1, -1, -1]
                    }
                ],
                Score: 3,
                HandsWon: 2
            }
        ],
        Trump: 0
    };
}]);
gameModule.filter('cardClasses', function(util) {
    return function(card) {
        if (card == -1) {
            return [];
        }
        return [util.getCardSuit(card), util.getCardNumber(card)];
    };
});
gameModule.factory('util', function() {
    var util = {};
    util.getCardSuit = function(card) {
        if (card < 6) { return 'heart'; }
        if (card < 12) { return 'spade'; }
        if (card < 18) { return 'diamond'; }
        return 'club';
    };
    util.getCardNumber = function(card) {
        switch (card % 6) {
        case 0: return 'nine';
        case 1: return 'ten';
        case 2: return 'jack';
        case 3: return 'queen';
        case 4: return 'king';
        case 5: return 'ace';
        }
    };
    return util;
});
gameModule.directive('hand', function() {
    return {
        restrict: 'E',
        templateUrl: 'game/hand.tpl.html',
        scope: {
            player: '='
        }
    };
});
