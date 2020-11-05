/*Refresh lobby users every 3 seconds*/
function refreshLobbyUsers() {
    $.ajax({
      url: "/lobby/lobbyUsers", success: function (result) {
        $("#card-container").html(users);
      }
    });
}

setInterval(refreshLobbyUsers, 3000)
/*Endlobby user refresh */

/*Invite player functionality */
function invitePlayer(userID) {
    $.post({
        url: `/lobby/invite/${userID}`, function (result, status) {
          //confim when good status
        }
    });
}

function checkInvite() {
    $.ajax({
        url: `/lobby/invite/check`, success: function (result) {
          
        }
    });
}

function acceptInvite() {
    $.ajax({
        url: `/lobby/invite/accept`, success: function (result) {
          
        }
    });
}

function denyInvite() {
    $.ajax({
        url: `/lobby/invite/deny`, success: function (result) {
          
        }
    });
}
/*End invite section */