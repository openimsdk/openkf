import { onUnmounted, Ref, ref } from 'vue';

// Return a count down tool
export const Counter = (duration = 60): [Ref<number>, () => void] => {
    let intervalTimer: ReturnType<typeof setInterval>;
    onUnmounted(() => {
        clearInterval(intervalTimer);
    });
    const countDown = ref(0);

    return [
        countDown,
        () => {
            countDown.value = duration;
            intervalTimer = setInterval(() => {
                if (countDown.value > 0) {
                    countDown.value -= 1;
                } else {
                    clearInterval(intervalTimer);
                    countDown.value = 0;
                }
            }, 1000);
        },
    ];
};

// Return a time of day
export const getTimeOfDay = (): string => {
    const currentTime = new Date();
    const currentHour = currentTime.getHours();

    if (currentHour >= 0 && currentHour < 12) {
        return 'Morning';
    } else if (currentHour >= 12 && currentHour < 18) {
        return 'Afternoon';
    } else {
        return 'Evening';
    }
}

// Return diff days between now and date 
export const getNowDiffDays = (date: string): number => {
    const diffTime = Math.abs(new Date(date).getTime() - new Date().getTime());
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));
    return diffDays;
}