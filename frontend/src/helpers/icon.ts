export function iconToUrlString(icon: string): string {
  return icon.trim().replace(/ /g, "_").replace(/'/g, "");
}

export function iconUrl(icon: string): string {
  const fileName = iconToUrlString(icon);
  return `https://oldschool.runescape.wiki/images/${encodeURIComponent(fileName)}`;
}
