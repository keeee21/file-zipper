export function downloadFile(downloadedFileName: string, downloadUrl: string) {
  const a = document.createElement('a');
  document.body.appendChild(a);
  a.download = downloadedFileName;
  a.href = downloadUrl;
  a.click();
  a.remove();
}
