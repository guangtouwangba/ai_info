# Upgrade or Switch: Do We Need a New Registry Architecture for the Internet of AI Agents?

**URL**: http://arxiv.org/abs/2506.12003v1

## 原始摘要

The emerging Internet of AI Agents challenges existing web infrastructure
designed for human-scale, reactive interactions. Unlike traditional web
resources, autonomous AI agents initiate actions, maintain persistent state,
spawn sub-agents, and negotiate directly with peers: demanding
millisecond-level discovery, instant credential revocation, and cryptographic
behavioral proofs that exceed current DNS/PKI capabilities. This paper analyzes
whether to upgrade existing infrastructure or implement purpose-built registry
architectures for autonomous agents. We identify critical failure points: DNS
propagation (24-48 hours vs. required milliseconds), certificate revocation
unable to scale to trillions of entities, and IPv4/IPv6 addressing inadequate
for agent-scale routing. We evaluate three approaches: (1) Upgrade paths, (2)
Switch options, (3) Hybrid registries. Drawing parallels to dialup-to-broadband
transitions, we find that agent requirements constitute qualitative, and not
incremental, changes. While upgrades offer compatibility and faster deployment,
clean-slate solutions provide better performance but require longer for
adoption. Our analysis suggests hybrid approaches will emerge, with centralized
registries for critical agents and federated meshes for specialized use cases.


## AI 摘要

本文探讨了AI智能体互联网对现有网络基础设施的挑战。传统网络架构无法满足AI代理的毫秒级发现、即时凭证撤销和行为加密证明等需求。研究分析了三种解决方案：升级现有设施、构建专用注册架构或混合方案。由于DNS传播延迟、证书撤销规模限制和IP地址不足等问题，AI代理的需求属于质变而非量变。虽然升级方案兼容性更好，但全新架构性能更优。最终可能形成混合模式，即关键代理使用集中注册，而特定场景采用联邦网状架构。研究类比了拨号到宽带的转型过程。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-16T03:25:59Z
- **目录日期**: 2025-06-16
