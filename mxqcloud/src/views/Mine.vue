<template>
    <Header/>
    <div class="mycard">
        <n-card :title="res" size="small">
            <n-form
                    ref="formRef"
                    :model="model"
                    label-placement="left"
                    :label-width="160"
                    :disabled="updateDisabled"
            >
                <n-form-item label="是否共享">
                    <n-switch v-model:value="model.isShare"/>
                </n-form-item>
                <n-form-item label="最大容量">
                    <n-input v-model:value="model.maxValue" placeholder="输入最大容量(G)"/>
                </n-form-item>
                <n-form-item v-show='isWindows' label="保存路径" v-model:value="model.savePath">
                    <n-input placeholder="输入上传保存的路径"/>
                </n-form-item>

                <div style="display: flex; justify-content: flex-end">
                    <n-space>
                        <n-button round type="primary" @click="handleModify">
                            修改
                        </n-button>
                        <n-button round type="primary" @click="handleSubmit">
                            确认
                        </n-button>
                    </n-space>

                </div>
            </n-form>
        </n-card>
    </div>

    <Footer/>
</template>

<script setup>
    import Header from '../components/Header.vue'
    import Footer from '../components/Footer.vue'
    import {ref} from 'vue'
    import {GetDevice} from '@/util/Device.js'

    let res = GetDevice();

    const model = ref({
        maxValue: 10,
        isShare: true,
        savePath: "",
    })
    const updateDisabled = ref(true);
    const formRef = ref(null);
    const handleModify = () => {
        updateDisabled.value = false;
    }
    const handleSubmit = () => {
        updateDisabled.value = true;
    }
    const isWindows = ref(res.includes("Windows"));

</script>

<style scoped lang="scss">
    .mycard {
        padding-top: $header-height;
        width: 90%;
        height: auto;
        margin-left: 5%;
        margin-right: 5%;
        margin-top: 10px;
    }
</style>