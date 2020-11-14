function checkGuesses() {
    $.ajax({
        url: `/game/guess/all`, type: "GET", 
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

var checkGuessesID = setInterval(checkGuesses, 5000);

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
            <line  y1="280" x1="75" y2="280" x2="225" stroke-width="3" stroke="#000" />
            <line  y1="35" x1="100" y2="280" x2="100" stroke-width="3" stroke="#000" />
            <line  y1="35" x1="100" y2="35" x2="200" stroke-width="3" stroke="#000" />
            <line  y1="35" x1="200" y2="70" x2="200"  stroke-width="3" stroke="#000" />
        </g>`;

    let head = `
            <circle cx="200" cy="95" r="25" stroke="black" stroke-width="3" fill="black" /> 
        `;

    let body = `
            <line  y1="120" x1="200" y2="215" x2="200" stroke-width="3" stroke="#000" />
        `;

    let lArm = `
            <line  y1="140" x1="200" y2="185" x2="170" stroke-width="3" stroke="#000" />
    `;

    let rArm = `
            <line  y1="140" x1="200" y2="185" x2="230" stroke-width="3" stroke="#000" />
    `;

    let lLeg = `
            <line  y1="215" x1="200" y2="255" x2="170" stroke-width="3" stroke="#000" />
    `;

    let rLeg = `
            <line  y1="215" x1="200" y2="255" x2="230" stroke-width="3" stroke="#000" />
    `;

    let svgArr = [gallow, head, body, lArm, rArm, lLeg, rLeg];
    
    let svg = "";
    for (i = 1; i <= count; i ++) {
        svg += svgArr[i-1]
    }

    $("#hangman-board").html(svg);
}

function endGame() {
    clearInterval(checkGuessesID);
    clearInterval(checkStatusID);
}