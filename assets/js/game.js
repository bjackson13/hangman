function checkIncorrectGuesses() {
    $.ajax({
        url: `/game/guess/incorrect`, type: "GET", success: function (result) {
            if(!result.responseJSON) {
                $("#guess-container").html(result);
                let numIncorrect = $("#incorrect-guesses").data("guessCount");
                drawHangman(numIncorrect)
            }
        },
        error: function(result) {
            console.log(result)
        }
    });   
}

var checkIncorrectID = setInterval(checkIncorrectGuesses, 5000);

function drawHangman(count) {

    let gallow = ``;

    let head = ``;

    let body = ``;

    let lArm = ``;

    let rArm = ``;

    let lLeg = ``;

    let rLeg = ``;

    let svgArr = [gallow, head, body, lArm, rArm, lLeg, rLeg];

    let svg = "";
    for (i = 0; i < count; i ++) {
        svg += svgArr[i]
    }

    $("hangman-board").html(svg);
}

function endGame() {
    clearInterval(checkIncorrectID);
}