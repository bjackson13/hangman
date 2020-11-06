/*Refresh lobby users every 3 seconds*/
function refreshLobbyUsers() {
    $.ajax({
      url: "/lobby/lobbyUsers", success: function (result) {
        $("#card-container").html(result);
      }
    });
}

setInterval(refreshLobbyUsers, 6000)
/*Endlobby user refresh */

/*Invite player functionality */
/*Set interval checks */
setInterval(checkInvite, 5000)

function invitePlayer(userID) {
    $.ajax({
        url: `/lobby/invite/user/${userID}`, type: "POST", success: function (result, status) {
          console.log(status);
          console.log(result);
        }
    });
}

function checkInvite() {
    $.ajax({
        url: `/lobby/invite/check`, success: function (result, status) {
          console.log(result)
          console.log(status)
        }
    });
}

function acceptInvite() {
    $.ajax({
        url: `/lobby/invite/accept`, type: "POST", success: function (result) {
          
        }
    });
}

function denyInvite() {
    $.ajax({
        url: `/lobby/invite/deny`, type: "POST", success: function (result) {
          
        }
    });
}
/*End invite section */