function makeGuess() {
    var guess = $("#keyboard").val()
    $.ajax({
        url: `/game/makeGuess`, 
        type: "POST",  
        data: {guess: guess}, 
        success: handleSuccess, 
        error: handleError
    });
    
    function handleSuccess(result) {
        if (result.error) {
            postMessage("danger", result.error)
        } else {
            $("#keyboard").val("")
            postMessage("success", result.success)
        }
    }

    function handleError(result) {
        console.log(result)
        postMessage("danger", result.responseJSON.error)
    }

    function postMessage(successOrDanger, message) {
        let html = `
            <div class="alert alert-${successOrDanger} alert-dismissible fade show" role="alert">
                ${message}
            <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                <span aria-hidden="true">&times;</span>
            </button>
            </div>
          `
    
        $("#keyboard-container").prepend(html);
    }

}