var gameServerModule = angular.module('GameServer', []);
gameServerModule.service('$GameServer', function() {
    var gameServer = function() {
        var _this = this;
        this.connect = function() {
            var gameId = window.location.hash.substring(7);
            var socketAddress = "ws://" + window.location.host + "/api/game/play/" + gameId + "/player/a/b/";
            socket = new WebSocket(socketAddress);
            socket.onmessage = function(m) {
                console.log(m);
            }
            socket.onopen = function(m) {
                console.log("Socket opened");
            }
            socket.onclose = function() {
                console.warn("Socket closed");
                _this.connect();
            }
        };
    };
    return new gameServer();
});