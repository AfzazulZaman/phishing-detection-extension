chrome.storage.local.get("detectionStatus", (data) => {
    if (data.detectionStatus) {
      const suspiciousKeywords = ["verify", "login", "bank", "password", "account"];
      const pageUrl = window.location.href;
  
      suspiciousKeywords.forEach((keyword) => {
        if (pageUrl.includes(keyword)) {
          alert("Warning: This site may be a phishing site!");
        }
      });
    }
  });
  