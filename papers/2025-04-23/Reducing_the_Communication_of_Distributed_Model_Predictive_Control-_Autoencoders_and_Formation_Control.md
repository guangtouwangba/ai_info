# Reducing the Communication of Distributed Model Predictive Control: Autoencoders and Formation Control

**URL**: http://arxiv.org/abs/2504.05223v2

## 原始摘要

Communication remains a key factor limiting the applicability of distributed
model predictive control (DMPC) in realistic settings, despite advances in
wireless communication. DMPC schemes can require an overwhelming amount of
information exchange between agents as the amount of data depends on the length
of the predication horizon, for which some applications require a significant
length to formally guarantee nominal asymptotic stability. This work aims to
provide an approach to reduce the communication effort of DMPC by reducing the
size of the communicated data between agents. Using an autoencoder, the
communicated data is reduced by the encoder part of the autoencoder prior to
communication and reconstructed by the decoder part upon reception within the
distributed optimization algorithm that constitutes the DMPC scheme. The choice
of a learning-based reduction method is motivated by structure inherent to the
data, which results from the data's connection to solutions of optimal control
problems. The approach is implemented and tested at the example of formation
control of differential-drive robots, which is challenging for
optimization-based control due to the robots' nonholonomic constraints, and
which is interesting due to the practical importance of mobile robotics. The
applicability of the proposed approach is presented first in form of a
simulative analysis showing that the resulting control performance yields a
satisfactory accuracy. In particular, the proposed approach outperforms the
canonical naive way to reduce communication by reducing the length of the
prediction horizon. Moreover, it is shown that numerical experiments conducted
on embedded computation hardware, with real distributed computation and
wireless communication, work well with the proposed way of reducing
communication even in practical scenarios in which full communication fails.


## AI 摘要

分布式模型预测控制(DMPC)在应用中面临通信量大的挑战，尤其当预测时域较长时。本研究提出利用自动编码器压缩通信数据：编码器在发送前压缩数据，解码器在接收后重建数据。这种方法特别适合处理与最优控制问题相关的结构化数据。以非完整约束的差速驱动机器人编队控制为例，仿真和嵌入式硬件实验表明，该方法在保证控制精度的同时显著降低通信量，优于传统的缩短预测时域方案，且在实际无线通信环境下表现良好，解决了完全通信失败的问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-23T18:02:10Z
- **目录日期**: 2025-04-23
