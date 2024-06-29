export const artworks = [
    {
        language: "go",
        name: "art_example",
    },
    {
        language: "go",
        name: "ruby_image",
    },
];
export function getArtWasm(a) {
    return eval(getArtWasmName(a));
}
export function getArtWasmName(a) {
    return `${a.language}_${a.name}`;
}
