export function simple() {}
/** description */
export function described() {}
/** @summary summary */
export function summarized() {}
export function async() {}
export function *generator() {}

export const arrowsimple = () => {}
/** arrow description */
export const arrowdescribed = () => {}
/** @summary arrow summary */
export const arrowsummarized = () => {}
export const arrowasync = () => {}

export function param(a: string): string {}
/**
 * description
 * @param a param description
 * @return return description
 */
export function describedparam(a: string): string {}
export function rest(...as: string[]): string {}
/**
 * description
 * @param as rest description
 * @return return description
 */
export function describedrest(...as: string[]): string {}
