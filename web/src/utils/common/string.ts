export const getAvatarString = (name: string): string => {
    return name.length > 2 ? name.slice(0, 2) : name.slice(0, 1);
}