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

本文提出了一种新型联邦学习框架FPoTT，用于解决车联网(IoV)中实时、可扩展且保护隐私的交通管理问题。该框架采用开源大语言模型(LLMs)，通过动态提示优化机制提升轨迹预测准确性，并采用双层联邦学习架构，结合轻量级边缘模型和云端LLMs。FPoTT还整合了基于Transformer的合成数据生成器，使用NGSIM格式数据增强训练。实验表明，使用Pythia-1B模型的FPoTT在真实数据上达到99.86%的预测准确率，为智能交通系统提供了安全、自适应且可扩展的解决方案，优于现有专有AI方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T00:01:52Z
- **目录日期**: 2025-05-05
