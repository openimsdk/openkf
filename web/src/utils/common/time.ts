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
