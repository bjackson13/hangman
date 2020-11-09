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
        url: `/game/guess/deny`, type: "GET", success: function (result) {
            console.log(result)
            $("#pending-guess-container").remove();
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
}