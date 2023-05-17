<template>
    <Header/>
    <div class="mycard">
        <n-card :title=model.Name size="small">
            <n-form
                    :model="model"
                    label-placement="left"
                    :label-width="160"
                    :disabled="updateDisabled"
            >
                <n-form-item label="是否共享">
                    <n-switch v-model:value="model.IsPublic"/>
                </n-form-item>
                <n-form-item label="最大容量">
                    <n-input-number
                            v-model:value="model.MaxSize"
                            placeholder="输入最大容量(G)"
                            :min="1"
                            :max="20"
                    />
                </n-form-item>
                <n-form-item v-show='isWindows' label="保存路径">
                    <n-input v-model:value="model.SavePath" placeholder="输入上传保存的路径"/>
                </n-form-item>

                <div style="display: flex; justify-content: flex-end">
                    <n-space>
                        <n-button size="tiny" round type="primary" @click="handleModify">
                            修改
                        </n-button>
                        <n-button size="tiny" round type="primary" @click="handleSubmit">
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
    import {ref, onBeforeMount} from 'vue'
    import {GetDevice} from '@/util/Device.js'
    import {AddDevice, GetDeviceByName} from '@/util/api.js'
    import {useNotification} from "naive-ui";

    const notification = useNotification();

    let res = GetDevice();

    const model = ref({
        Name: res,
        MaxSize: 10,
        IsPublic: true,
        SavePath: "",
        IsServer: false
    })


    const updateDisabled = ref(true);
    const handleModify = () => {
        updateDisabled.value = false;
    }

    const isWindows = ref(res === "Windows PC");

    onBeforeMount(() => {
        GetDeviceByName(res).then((ress) => {
            if (ress.code === 200) {
                model.value = ress.data;
            }
        });
    });

    const handleSubmit = async () => {
        if (isWindows.value) {
            model.value.IsServer = true
        }
        if (isWindows.value && model.value.SavePath.length === 0) {
            notification.error(
                {
                    content: "修改失败，服务端保存路径无效",
                    duration: 3000,
                }
            )
            return
        }
        updateDisabled.value = true;
        await AddDevice(model.value).then((res) => {
            if (res.code === 200) {
                notification.success(
                    {
                        content: "修改成功",
                        duration: 3000,
                    }
                )
            } else {
                notification.error(
                    {
                        content: "修改失败",
                        duration: 3000,
                    }
                )
            }
        })
    }


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