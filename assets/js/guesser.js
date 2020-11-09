function makeGuess() {
    guess = $("#keyboard").val()
    $.ajax({
        url: `/game/makeGuess`, type: "POST",  data: {guess: guess}, success: function (result) {
            let html = `
            <div class="alert alert-success alert-dismissible fade show" role="alert">
                ${result.success}
            <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                <span aria-hidden="true">&times;</span>
            </button>
            </div>
          `
          $("#keyboard-container").prepend(html);

            $("#keyboard").val("")
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
          $("#keyboard-container").prepend(html);
        }
    });
}