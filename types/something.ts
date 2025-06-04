export enum Code {
    NotFound
}

export const HTTPCode = (c: Code) => {
    return [404][c]
}
export const CustomCode = (c: Code) => {
    return [404][c]
}
