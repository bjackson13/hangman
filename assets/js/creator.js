/*Check guesses*/
function checkGuess() {
    $.ajax({
        url: `/game/guess`, type: "GET", success: function (result) {
            if(!result.responseJSON) {
                $("#pending-stage").html(result);
            }
        },
        error: function(result) {
            console.log(result)
        }
    });    
}

setInterval(checkGuess, 6000)
/*End checking guesses */

function denyGuess() {
    $.ajax({
        url: `/game/guess/deny`, type: "GET", 
        success: function (result) {
            $("#pending-guess-container").remove();
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
          $("#pending-guess-container").prepend(html);
        }
    });    
}

function acceptGuess() {
    let indexes = [];

    $.ajax({
        url: `/game/guess/accept`, type: "POST", data: {indexes: indexes},
        success: function (result) {
            $("#pending-guess-container").remove();
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
          $("#pending-guess-container").prepend(html);
        }
    });  

}

function submitWord() {
    let length = $("#word-length").val();
    if (!(length <= 0 || length > 15 )) {
        $.ajax({
            url: `/game/word/create`, type: "POST", data: {length: length}, success: function (result) {
                console.log(result)
                $("#word-creation-modal, .modal-backdrop").remove();
            },
            error: function(result) {
                console.log(result)
                let html = `
                <div class="alert alert-danger alert-dismissible fade show" role="alert">
                    ${result.error}
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
                </div>
            `
            $("#pending-guess-container").prepend(html);
            }
        });
    } else {
        $("#word-length").css("border", "3px solid red");
    }
}