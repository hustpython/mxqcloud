<template>
    <n-layout-footer
            position="absolute"
            style="height: auto; padding: 0;"

    >
        <n-space class="footer" justify="space-between">
            <router-link :to="{ path: '/' }">
                <n-icon size="30"
                        :color=iconColor[0]>
                    <CloudUploadSharp/>
                    <div :style="color=iconColor[0]"
                         class="footer-font">
                        上传
                    </div>
                </n-icon>
            </router-link>
            <router-link :to="{ path: '/download' }">
                <n-icon size="30" :color="color=iconColor[1]">
                    <CloudDownloadRound/>
                    <div
                            :style="color=iconColor[1]"
                            class="footer-font">
                        下载
                    </div>
                </n-icon>
            </router-link>
            <router-link :to="{ path: '/chat' }">
                <n-icon size="30" :color="color=iconColor[2]">
                    <WechatSharp/>
                    <div :style="color=iconColor[2]"
                         class="footer-font">
                        聊天
                    </div>
                </n-icon>
            </router-link>

            <router-link :to="{ path: '/mine' }">
                <n-icon size="30" :color="color=iconColor[3]">
                    <PersonRound/>
                    <div :style="color=iconColor[3]"
                         class="footer-font">
                        我的
                    </div>
                </n-icon>
            </router-link>

        </n-space>
    </n-layout-footer>


</template>

<script setup>
    import {CloudUploadSharp, CloudDownloadRound, WechatSharp, PersonRound} from "@vicons/material";
    import {ref} from 'vue'
    import {useRouter} from "vue-router";

    const router = useRouter();

    const iconColor = ref(['', '', '', ''])
    const routeNameArry = ['upload', 'download', 'chat', 'mine']
    routeNameArry.forEach((k, i) => {
        if (router.currentRoute.value.name === k) {
            iconColor.value[i] = 'green'
        }
    })
</script>

<style scoped lang="scss">
    .footer {
        margin-left: 30px;
        margin-right: 30px;

        .footer-font {
            font-size: 12px;
            font-style: normal;
            font-family: 'Microsoft YaHei';
            text-align: center;
            margin-top: -10px;
        }
    }
</style>