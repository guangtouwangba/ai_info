# Open-Source LLM-Driven Federated Transformer for Predictive IoV Management

**URL**: http://arxiv.org/abs/2505.00651v2

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

本文提出了一种名为FPoTT（联邦优化提示交通转换器）的新框架，旨在解决车联网（IoV）中可扩展性、实时性和隐私保护的挑战。FPoTT利用开源大型语言模型（LLM），通过动态提示优化机制提升轨迹预测准确性。该框架采用双层联邦学习架构，结合轻量级边缘模型和云端LLM，并引入基于Transformer的合成数据生成器来增强训练。实验表明，FPoTT在真实数据上实现了99.86%的预测准确率，同时在合成数据集上也表现优异。这一成果为车联网管理提供了安全、自适应且可扩展的解决方案，有望替代现有专有方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T19:01:27Z
- **目录日期**: 2025-05-14
