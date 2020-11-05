function refreshLobbyUsers() {
    $.ajax({
      url: "/lobby/lobbyUsers", success: function (result) {
        appendLobbyUsers(result);
      }
    });
}

function appendLobbyUsers(users) {
    $("#card-container").html(users);
}

setInterval(refreshLobbyUsers, 5000)