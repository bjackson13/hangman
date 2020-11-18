var LAST_UPDATED_CHAT = 0;

function getNewMessages() {
    $.ajax({
        url: `/chat/since/${LAST_UPDATED_CHAT}`, type: "GET", 
        success: function (result) {
            if(!result.responseJSON) {
                $("#messages").append(result);
                LAST_UPDATED_CHAT = Math.floor(Date.now() / 1000);
            }
        },
        error: function(result) {
            console.log(result)
        }
    });
}

function sendMessage() {
    let message = $("#message-box").val().trim();
    if (message.length) {
        $.ajax({
            url: `/chat/`, 
            data: {message: message},
            type: "POST", 
            success: function (result) {
                $("#message-box").val("");
                console.log(result)
            },
            error: function(result) {
                console.log(result)
            }
        });
    }
}

function getAllMessages() {
    $.ajax({
        url: `/chat/`, type: "GET", 
        success: function (result) {
            if(!result.responseJSON) {
                $("#messages").html(result);
                LAST_UPDATED_CHAT = Math.floor(Date.now() / 1000);
                setInterval(getNewMessages, 3000)
            }
        },
        error: function(result) {
            console.log(result)
        }
    });
}

$(document).ready(function() {
    getAllMessages();
});

