export const getAvatarString = (name: string): string => {
    return name.length > 2 ? name.slice(0, 2) : name.slice(0, 1);
}

// To capitail the first
export const capitalizeFirstLetter = (str: string): string => {
    return str.replace(/\b[a-z]/g, (char) => char.toUpperCase());
}

// Change all letters to upper case
export const toUpperCase = (str: string): string => {
    return str.toUpperCase();
}