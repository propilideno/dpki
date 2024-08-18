<template>
  <q-dialog
    ref="dialogRef"
    @hide="onDialogHide"
    @shake="onCloseDialog"
    :position="$q.screen.gt.xs ? 'standard' : 'bottom'"
  >
    <q-card class="medium-dialog">
      <q-form @submit="onSubmit">
        <q-card-section class="row no-wrap">
          <div class="full-width-flex">
            <div class="text-h6 text-grey-9 flex items-center full-height">
              Sign Message
            </div>
          </div>
        </q-card-section>
        <q-separator />
        <q-card-section>
          <div class="row q-col-gutter-md">
            <div class="col-12">
              <q-input
                v-model="signParams.message"
                label="Message *"
                :rules="rules.message"
                filled
              />
            </div>
            <div class="col-12">
              <q-input
                v-model="signParams.private_key"
                label="Private Key *"
                :rules="rules.private_key"
                filled
                autogrow
              />
            </div>
            <div class="col-12">
              <q-input
                v-model="sign"
                label="Signature"
                autogrow
                readonly
                filled
                input-class="text-caption"
              >
                <template #append>
                  <q-btn
                    @click="copyKeyToClipboard(sign)"
                    icon="content_copy"
                    size="0.5em"
                    round
                  >
                    <q-tooltip class="text-subtitle2">
                      Copy Signature
                    </q-tooltip>
                  </q-btn>
                </template>
              </q-input>
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn
            @click="onCloseDialog"
            label="Close"
            color="grey-9"
            padding="sm md"
            flat
          />
          <q-btn
            type="submit"
            label="Generate"
            color="primary"
            text-color="white"
            padding="sm md"
            :loading="loading"
          />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
</template>

<script>
import { defineComponent, ref, reactive } from "vue";
import { useDialogPluginComponent, copyToClipboard, useQuasar } from "quasar";
import cloneDeep from "lodash.clonedeep";
import { api } from "src/boot/axios";
import { requiredRule } from "src/utils/rules";

const DEFAULT_SIGN_PARAMS = {
  message: "",
  private_key: "",
};

const useRules = () => {
  return {
    message: [requiredRule()],
    private_key: [requiredRule()],
  };
};

export default defineComponent({
  name: "SignMessageDialog",
});
</script>

<script setup>
/* const props = */ defineProps({});

/* const emit = */ defineEmits([...useDialogPluginComponent.emits]);

const {
  dialogRef,
  onDialogHide,
  onDialogOK: onOKClick,
  onDialogCancel: onCancelClick,
} = useDialogPluginComponent();

const onCloseDialog = () => {
  onCancelClick();
};

const $q = useQuasar();

const rules = reactive(useRules());

const sign = ref("");
const signParams = ref(cloneDeep(DEFAULT_SIGN_PARAMS));
const loading = ref(false);

const copyKeyToClipboard = (value) => {
  copyToClipboard(value)
    .then(() => {
      $q.notify({
        type: "positive",
        message: "Copied!",
      });
    })
    .catch(() => {
      $q.notify({
        type: "alert",
        message: "Error on Copy.",
      });
    });
};

// TODO: adicionar rules

const onSubmit = async () => {
  sign.value = "";

  try {
    loading.value = true;

    const { data } = await api.get("sign", {
      params: signParams.value,
    });

    sign.value = data.sign;
  } finally {
    loading.value = false;
  }
};
</script>
