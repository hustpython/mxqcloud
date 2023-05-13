<script setup>
    const handleUpload = () => {
        navigator.camera.getPicture(
            function (imageURI) {
                // 处理获取的照片
                alert('Image URI: ' + imageURI);
            },
            function (message) {
                // 处理错误
                alert('Failed to get picture: ' + message);
            },
            {
                sourceType: Camera.PictureSourceType.PHOTOLIBRARY,
                destinationType: Camera.DestinationType.FILE_URI,
                quality: 100,
                encodingType: Camera.EncodingType.JPEG,
                correctOrientation: true,
                mediaType: Camera.MediaType.PICTURE
            }
        );
    }
</script>

<template>
    <n-button @click="handleUpload">
        上传
    </n-button>
</template>

<style scoped>

</style>
