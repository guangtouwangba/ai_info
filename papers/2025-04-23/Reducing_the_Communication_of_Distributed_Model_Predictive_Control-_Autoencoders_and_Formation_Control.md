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

分布式模型预测控制(DMPC)在实际应用中面临通信量过大的问题。本研究提出利用自编码器压缩传输数据，在保持预测时域长度的同时减少通信负担。该方法通过编码器压缩数据、解码器重建数据，特别适用于具有最优控制问题结构的数据。以差速驱动机器人编队控制为例进行验证，结果表明：相比直接缩短预测时域的传统方法，新方法在保证控制精度的同时显著降低通信量。仿真和嵌入式硬件实验均证实，该方案在实际通信受限场景中表现优于全通信方案，具有实用价值。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-23T20:02:24Z
- **目录日期**: 2025-04-23
