# Schemas
## URN Schema

Below is an experiment in defining a URN that fully defines an embedding model by providing all necessary information to
create an embedding adapter.

`urn:embedding:<schema-version>/<provider>/<vendor>/<model>:<model-version>/<modality>?<config-params>`

**Components:**

- `schema-version`: Version of the URN schema (e.g., 1)
- `provider`: The provider that operates the model.
- `vendor`: Provider of the embedding model (e.g., openai, google, huggingface)
- `model`: Name of the specific model
- `model-version`: Version or checkpoint of the model
- `modality`: Specific modality or mode of the model (e.g., text, image, audio, multi)
- `config-params`: Optional configuration parameters (excluding API keys)

**Examples:**

- `urn:embedding:1/openai/openai/text-embedding-ada-002:1/text?max_batch_size=100`
- `urn:embedding:1/googlegenerativeai/google/multimodal-bert:2/multi`
- `urn:embedding:1//huggingface/clip-vit-base-patch32:1/image?image_size=224`
- `urn:embedding:1/openai/whisper:1/audio?sample_rate=16000`