<template>
    <Header/>
    <div class="download">
        下载
    </div>
    <Footer/>
</template>

<script setup>
    import Header from '@/components/Header.vue'
    import Footer from '@/components/Footer.vue'
</script>

<style scoped lang="scss">
.download {
    padding-top: $header-height;
    width: 100%;
    height: auto;
}
</style>