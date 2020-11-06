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
        url: `/lobby/invite/check`, success: function (result, status, xhr, datatype) {
            if (datatype == "json") {
                console.log(result)
                console.log(status)
            } else {
                console.log(result)
                $("#modal-container").html(result);
                $('#invite-modal').modal("show");
            }
        }
    });
}

function acceptInvite(inviterID) {
    $.ajax({
        url: `/lobby/invite/accept`, data: {inviterID: inviterID}, type: "POST", success: function (result) {
          console.log(result);
        }
    });
}

function denyInvite() {
    $.ajax({
        url: `/lobby/invite/deny`, type: "POST", success: function (result) {
            console.log(result);
            console.log(status);
        }
    });
}
/*End invite section */