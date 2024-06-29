export interface Artwork {
    code: string;
    name: string;
    title: string;
    description: string;
}
  
export const artworks: Artwork[] = [
    { 
        code: "go", 
        name: "art_example", 
        title: "memo",
        description: "memoです"
    },
    { 
        code: "go", 
        name: "ruby_image", 
        title: "memo",
        description: "memoです"
    },
];

export function getArtWasm(a: Artwork) {
    return eval(getArtWasmName(a))
}

export function getArtWasmName(a: Artwork) {
    return `${a.code}_${a.name}`
}