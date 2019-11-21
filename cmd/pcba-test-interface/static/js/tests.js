async function runTests() {
  $("#run-tests-button").html("Running Tests").prop("disabled", true)
  const tests  = [
    {name: "RTC", path: "/api/test-rtc", resultElemetnID: "rtc-results"},
    {name: "ATtiny", path: "/api/test-attiny", resultElemetnID: "attiny-results"},
    {name: "USB", path: "/api/test-usb", resultElemetnID: "usb-results"},
  ]
  for (const test of tests) {
    $("#"+test.resultElemetnID).html("")
  }
  for (const test of tests) {
    console.log("testing " + test.name);
    var result = await apiGetJSON(test.path);
    console.log(result);
    if (result.Failed == null) {
      $("#"+test.resultElemetnID).html("✓")
    } else {
      $("#"+test.resultElemetnID).html(result.Failed.join(", "))
    }
  }
  $("#run-tests-button").html("Run Tests").prop("disabled", false)
}

async function playSound() {
  $("#play-sound-button").html("Playing Sound").prop("disabled", true)
  await apiGetJSON("/api/test-speaker");
  $("#play-sound-button").html("Play Sound").prop("disabled", false)
}


function apiGetJSON(url) {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open("GET", url, true);
    xhr.setRequestHeader("Authorization", "Basic "+btoa("admin:feathers"))
    xhr.onload = () => {
      if (200 <= xhr.status && xhr.status < 300) {
        resolve(JSON.parse(xhr.responseText));
      } else {
        reject(xhr)
      }
    }
    xhr.onerror = () => reject(xhr);
    xhr.send();
  });
}
