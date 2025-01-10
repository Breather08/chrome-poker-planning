chrome.runtime.onInstalled.addListener(() => {
  console.log("Chrome extension installed.");
});

chrome.sidePanel
  .setPanelBehavior({ openPanelOnActionClick: true })
  .catch((error) => console.error(error));
