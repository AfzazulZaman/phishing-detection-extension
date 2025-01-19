chrome.tabs.onUpdated.addListener((tabId, changeInfo, tab) => {
    chrome.storage.local.get("detectionStatus", (data) => {
      if (data.detectionStatus && changeInfo.url) {
        checkPhishingSite(changeInfo.url);
      }
    });
  });
  
  function checkPhishingSite(url) {
    const phishingIndicators = ["login", "verify", "update", "bank", "password", "account"];
    const isPhishing = phishingIndicators.some(keyword => url.includes(keyword));
  
    if (isPhishing) {
      chrome.scripting.executeScript({
        target: { tabId: tabId },
        func: showPhishingWarning
      });
    }
  }
  
  function showPhishingWarning() {
    const warningDiv = document.createElement('div');
    warningDiv.id = 'fullscreenWarning';
    warningDiv.innerHTML = `
      <h1>Warning: This is a Phishing Site!</h1>
      <p>Do not proceed. It's unsafe to share any personal information.</p>
    `;
    document.body.appendChild(warningDiv);
    warningDiv.style.display = 'flex';
  }
  