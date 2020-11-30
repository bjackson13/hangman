/*Refresh lobby users every 3 seconds*/
function refreshLobbyUsers() {
    $.ajax({
      url: "/lobby/lobbyUsers", 
      success: function (result, status) {
        $("#card-container").html(result);
      },
      error: function(result) {
        if (result.status == 302) {
            window.location.href = result.responseJSON.url
        }
      }
    });
}

/*Endlobby user refresh */

/*Invite player functionality */
function invitePlayer(userID) {
    $.ajax({
        url: `/lobby/invite/user/${userID}`, type: "POST", success: function (result) {
            console.log(result)
          let html = `
            <div class="alert alert-success alert-dismissible fade show" role="alert">
                ${result.success}
            <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                <span aria-hidden="true">&times;</span>
            </button>
            </div>
          `
            $("body").prepend(html);
        },
        error: function(result) {
            let html = `
            <div class="alert alert-danger alert-dismissible fade show" role="alert">
                ${result.error}
            <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                <span aria-hidden="true">&times;</span>
            </button>
            </div>
          `
          $("body").prepend(html);
        }
    });
}

function checkInvite() {
    $.ajax({
        url: `/lobby/invite/check`, success: function (result) {
            if (!result.responseJSON) {
                $("#modal-container").html(result);
                $('#invite-modal').modal("show");
            }
        }
    });
}

function acceptInvite(inviterID) {
    $.ajax({
        url: `/lobby/invite/accept`, data: {inviterID: inviterID}, type: "POST", success: function (result) {
          window.location.href = "/game"
        }
    });
}

function denyInvite() {
    $.ajax({
        url: `/lobby/invite/deny`, type: "POST", success: function (result) {
            $("#invite-modal, .modal-backdrop").remove();
        }
    });
}
/*End invite section */

$( document ).ready(function() {
    /*Set interval checks */
    setInterval(refreshLobbyUsers, 5000);
    setInterval(checkInvite, 5000);
});
