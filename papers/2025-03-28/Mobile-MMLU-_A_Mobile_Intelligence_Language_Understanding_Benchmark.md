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

研究人员开发了Mobile-MMLU基准数据集，用于评估移动设备上大型语言模型（LLM）的性能。该数据集包含16,186个问题，涵盖80个移动相关领域，模拟真实移动场景如食谱推荐、旅行规划等日常任务。其进阶版Mobile-MMLU-Pro难度显著提高。该基准特别关注移动环境的关键指标：推理延迟、能耗、内存占用和响应质量，同时强调隐私保护和个性化适应能力。通过提供标准化评估框架，Mobile-MMLU旨在推动移动端LLM的优化发展，解决移动设备在存储和计算资源上的限制问题。数据集和代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-28T02:27:23Z
- **目录日期**: 2025-03-28
