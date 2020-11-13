function checkIncorrectGuesses() {
    $.ajax({
        url: `/game/guess/incorrect`, type: "GET", 
        success: function (result) {
            if(!result.responseJSON) {
                $("#guess-container").html(result);
                let numIncorrect = $("#incorrect-guesses").data("guesscount");
                drawHangman(numIncorrect);
            }
        },
        error: function(result) {
            console.log(result)
            window.location.href = "/lobby";
        }
    });   
}

var checkIncorrectID = setInterval(checkIncorrectGuesses, 5000);

function checkGameStatus() {
    $.ajax({
        url: `/game/status`, type: "GET", 
        success: function (result) {
            if(!result.responseJSON) {
                $("body").append(result);
                $("#endgame-modal").modal("show");
            }
        },
        error: function(result) {
            console.log(result)
        }
    }); 
}

var checkStatusID = setInterval(checkGameStatus, 5000);

function drawHangman(count) {

    let gallow = `
        <g>
            <line y1="250" x1="75" y2="250" x2="225" stroke-width="3" stroke="#000" />
            <line  y1="35" x1="100" y2="250" x2="100" stroke-width="3" stroke="#000" />
            <line  y1="35" x1="100" y2="35" x2="200" stroke-width="3" stroke="#000" />
            <line  y1="35" x1="200" y2="85" x2="200"  stroke-width="3" stroke="#000" />
        </g>`;

    let head = `
    
    `;

    let body = ``;

    let lArm = ``;

    let rArm = ``;

    let lLeg = ``;

    let rLeg = ``;

    let svgArr = [gallow, head, body, lArm, rArm, lLeg, rLeg];
    
    let svg = "";
    for (i = 1; i <= count; i ++) {
        svg += svgArr[i-1]
    }

    $("#hangman-board").html(svg);
}

function endGame() {
    clearInterval(checkIncorrectID);
    clearInterval(checkStatusID);
}