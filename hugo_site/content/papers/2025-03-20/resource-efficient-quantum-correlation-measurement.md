---
title: "Resource-Efficient Quantum Correlation Measurement: A Multicopy Neural Network Approach for Practical Applications"
date: 2025-03-20T13:02:00+00:00
source_url: "http://arxiv.org/abs/2411.05745v2"
categories: ["2025-03-20"]
tags: ["Measurement", "Approach", "Practical", "Quantum", "Correlation"]
summary: "传统量子态层析（QST）方法需要重建密度矩阵，资源需求随系统规模指数增长。本文提出一种结合多拷贝测量和人工神经网络（ANN）的策略，将测量需求减少67%。实验在IBMQ处理器上使用transmon量子比特，成功测量了受退极化通道和振幅阻尼通道影响的两量子比特纠缠态。通过最大似然方法抑制噪声，并利用SHAP分析优化投影集训练ANN估计纠缠和非局域性。该方法简化了多拷贝测量，提高了噪声鲁棒性，无需复杂非线性方程分析，显著推动了AI辅助量子测量的实际应用。"
---

# Resource-Efficient Quantum Correlation Measurement: A Multicopy Neural Network Approach for Practical Applications

**原始链接**: [查看原文](http://arxiv.org/abs/2411.05745v2)

## 原始摘要

Measuring complex properties in quantum systems, such as measures of quantum
entanglement and Bell nonlocality, is inherently challenging. Traditional
methods, like quantum state tomography (QST), necessitate a full reconstruction
of the density matrix for a given system and demand resources that scale
exponentially with system size. We propose an alternative strategy that reduces
the required information by combining multicopy measurements with artificial
neural networks (ANNs), resulting in a 67\% reduction in measurement
requirements compared to QST. We have successfully measured two-qubit quantum
correlations of Bell states subjected to a depolarizing channel (resulting in
Werner states) and an amplitude damping channel (leading to Horodecki states)
using the multicopy approach on real quantum hardware. Our experiments,
conducted with transmon qubits on IBMQ processors, quantified the violation of
Bell's inequality and the negativity of two-qubit entangled states. We compared
these results with those obtained from the standard QST approach and applied a
maximum likelihood method to mitigate noise. We trained ANNs to estimate
entanglement and nonlocality measures using optimized sets of projections
identified through Shapley's (SHAP) analysis for the Werner and Horodecki
states. The ANN output, based on this reduced set of projections, aligns well
with expected values and enhances noise robustness. This approach simplifies
and improves the error-robustness of multicopy measurements, eliminating the
need for complex nonlinear equation analysis. It represents a significant
advancement in AI-assisted quantum measurements, making practical
implementation on current quantum hardware more feasible.

## AI 摘要

传统量子态层析（QST）方法需要重建密度矩阵，资源需求随系统规模指数增长。本文提出一种结合多拷贝测量和人工神经网络（ANN）的策略，将测量需求减少67%。实验在IBMQ处理器上使用transmon量子比特，成功测量了受退极化通道和振幅阻尼通道影响的两量子比特纠缠态。通过最大似然方法抑制噪声，并利用SHAP分析优化投影集训练ANN估计纠缠和非局域性。该方法简化了多拷贝测量，提高了噪声鲁棒性，无需复杂非线性方程分析，显著推动了AI辅助量子测量的实际应用。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-20T13:02:45+08:00
- **目录日期**: 2025-03-20
