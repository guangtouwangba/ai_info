# Mobile-MMLU: A Mobile Intelligence Language Understanding Benchmark

**URL**: http://arxiv.org/abs/2503.20786v1

## 原始摘要

Rapid advancements in large language models (LLMs) have increased interest in
deploying them on mobile devices for on-device AI applications. Mobile users
interact differently with LLMs compared to desktop users, creating unique
expectations and data biases. Current benchmark datasets primarily target at
server and desktop environments, and there is a notable lack of extensive
datasets specifically designed for mobile contexts. Additionally, mobile
devices face strict limitations in storage and computing resources,
constraining model size and capabilities, thus requiring optimized efficiency
and prioritized knowledge. To address these challenges, we introduce
Mobile-MMLU, a large-scale benchmark dataset tailored for mobile intelligence.
It consists of 16,186 questions across 80 mobile-related fields, designed to
evaluate LLM performance in realistic mobile scenarios. A challenging subset,
Mobile-MMLU-Pro, provides advanced evaluation similar in size to MMLU-Pro but
significantly more difficult than our standard full set. Both benchmarks use
multiple-choice, order-invariant questions focused on practical mobile
interactions, such as recipe suggestions, travel planning, and essential daily
tasks. The dataset emphasizes critical mobile-specific metrics like inference
latency, energy consumption, memory usage, and response quality, offering
comprehensive insights into model performance under mobile constraints.
Moreover, it prioritizes privacy and adaptability, assessing models' ability to
perform on-device processing, maintain user privacy, and adapt to personalized
usage patterns. Mobile-MMLU family offers a standardized framework for
developing and comparing mobile-optimized LLMs, enabling advancements in
productivity and decision-making within mobile computing environments. Our code
and data are available at: https://github.com/VILA-Lab/Mobile-MMLU.


## AI 摘要

研究人员开发了Mobile-MMLU基准数据集，专门用于评估移动设备上的大语言模型（LLM）性能。该数据集包含16,186个问题，涵盖80个移动相关领域，并设有更具挑战性的Mobile-MMLU-Pro子集。它重点关注移动场景下的实际需求（如食谱推荐、旅行规划等），并评估推理延迟、能耗、内存使用等关键指标。该基准强调隐私保护和个性化适配能力，为移动优化LLM的开发与比较提供了标准化框架，推动移动AI应用的发展。数据集和代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-27T12:01:15Z
- **目录日期**: 2025-03-27
