# Open-Source LLM-Driven Federated Transformer for Predictive IoV Management

**URL**: http://arxiv.org/abs/2505.00651v1

## 原始摘要

The proliferation of connected vehicles within the Internet of Vehicles (IoV)
ecosystem presents critical challenges in ensuring scalable, real-time, and
privacy-preserving traffic management. Existing centralized IoV solutions often
suffer from high latency, limited scalability, and reliance on proprietary
Artificial Intelligence (AI) models, creating significant barriers to
widespread deployment, particularly in dynamic and privacy-sensitive
environments. Meanwhile, integrating Large Language Models (LLMs) in vehicular
systems remains underexplored, especially concerning prompt optimization and
effective utilization in federated contexts. To address these challenges, we
propose the Federated Prompt-Optimized Traffic Transformer (FPoTT), a novel
framework that leverages open-source LLMs for predictive IoV management. FPoTT
introduces a dynamic prompt optimization mechanism that iteratively refines
textual prompts to enhance trajectory prediction. The architecture employs a
dual-layer federated learning paradigm, combining lightweight edge models for
real-time inference with cloud-based LLMs to retain global intelligence. A
Transformer-driven synthetic data generator is incorporated to augment training
with diverse, high-fidelity traffic scenarios in the Next Generation Simulation
(NGSIM) format. Extensive evaluations demonstrate that FPoTT, utilizing
EleutherAI Pythia-1B, achieves 99.86% prediction accuracy on real-world data
while maintaining high performance on synthetic datasets. These results
underscore the potential of open-source LLMs in enabling secure, adaptive, and
scalable IoV management, offering a promising alternative to proprietary
solutions in smart mobility ecosystems.


## AI 摘要

本文提出了一种新型联邦学习框架FPoTT，用于解决车联网（IoV）中的实时、可扩展和隐私保护的交通管理问题。FPoTT采用开源大语言模型（LLMs），通过动态提示优化机制提升轨迹预测精度。该框架采用双层联邦学习架构，结合轻量级边缘模型和云端LLMs，并利用Transformer生成合成数据增强训练。实验表明，FPoTT在真实数据上达到99.86%的预测准确率，同时保持对合成数据的高性能。这一方案为智能交通系统提供了一种安全、自适应且可扩展的替代方案，克服了现有集中式解决方案的延迟和扩展性限制。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-02T18:01:42Z
- **目录日期**: 2025-05-02
