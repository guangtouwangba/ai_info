# FLIP Reasoning Challenge

**URL**: http://arxiv.org/abs/2504.12256v1

## 原始摘要

Over the past years, advances in artificial intelligence (AI) have
demonstrated how AI can solve many perception and generation tasks, such as
image classification and text writing, yet reasoning remains a challenge. This
paper introduces the FLIP dataset, a benchmark for evaluating AI reasoning
capabilities based on human verification tasks on the Idena blockchain. FLIP
challenges present users with two orderings of 4 images, requiring them to
identify the logically coherent one. By emphasizing sequential reasoning,
visual storytelling, and common sense, FLIP provides a unique testbed for
multimodal AI systems. Our experiments evaluate state-of-the-art models,
leveraging both vision-language models (VLMs) and large language models (LLMs).
Results reveal that even the best open-sourced and closed-sourced models
achieve maximum accuracies of 75.5% and 77.9%, respectively, in zero-shot
settings, compared to human performance of 95.3%. Captioning models aid
reasoning models by providing text descriptions of images, yielding better
results than when using the raw images directly, 69.6% vs. 75.2% for Gemini 1.5
Pro. Combining the predictions from 15 models in an ensemble increases the
accuracy to 85.2%. These findings highlight the limitations of existing
reasoning models and the need for robust multimodal benchmarks like FLIP. The
full codebase and dataset will be available at
https://github.com/aplesner/FLIP-Reasoning-Challenge.


## AI 摘要

这篇论文介绍了FLIP数据集，用于评估AI在基于Idena区块链人类验证任务中的推理能力。FLIP要求用户从两组4张图片序列中识别逻辑连贯的排序，测试顺序推理、视觉叙事和常识理解。实验显示，最佳开源和闭源模型在零样本设置下的准确率分别为75.5%和77.9%，远低于人类的95.3%。使用图像描述模型辅助推理可提升性能（Gemini 1.5 Pro从69.6%升至75.2%），而15个模型的集成方法将准确率提高到85.2%。结果表明现有推理模型存在局限，凸显了FLIP等多模态基准的重要性。代码和数据集已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-17T23:01:19Z
- **目录日期**: 2025-04-17
