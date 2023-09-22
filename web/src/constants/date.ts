import dayjs from 'dayjs';

export const LAST_7_DAYS = [
    // today - 7
    dayjs().subtract(5, 'day').format('YYYY-MM-DD'),
    dayjs().subtract(-1, 'day').format('YYYY-MM-DD'), // today
];

export const LAST_30_DAYS = [
    dayjs().subtract(30, 'day').format('YYYY-MM-DD'),
    dayjs().subtract(1, 'day').format('YYYY-MM-DD'),
];
