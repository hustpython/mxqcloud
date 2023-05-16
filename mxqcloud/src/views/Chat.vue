<template>
    <Header/>
    <div class="chat">
        聊天
    </div>
    <Footer/>
</template>

<script setup>
    import Header from '@/components/Header.vue'
    import Footer from '@/components/Footer.vue'
</script>

<style scoped lang="scss">
    .chat {
        padding-top: $header-height;
        width: 100%;
        height: auto;
    }
</style>