document.addEventListener("DOMContentLoaded", () => {
    const toggle = document.getElementById("toggle");
    const statusElement = document.getElementById("status");
  
    // Load the detection status
    chrome.storage.local.get("detectionStatus", (data) => {
      toggle.checked = data.detectionStatus || false;
      statusElement.textContent = toggle.checked ? "Detection is ON" : "Detection is OFF";
    });
  
    // Toggle functionality
    toggle.addEventListener("change", (e) => {
      const isEnabled = e.target.checked;
      chrome.storage.local.set({ detectionStatus: isEnabled });
      statusElement.textContent = isEnabled ? "Detection is ON" : "Detection is OFF";
    });
  });
  