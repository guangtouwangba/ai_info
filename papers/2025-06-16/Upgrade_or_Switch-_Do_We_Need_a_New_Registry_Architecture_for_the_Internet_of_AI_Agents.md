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

本文探讨了现有网络基础设施在应对新兴AI代理互联网（Internet of AI Agents）时面临的挑战。传统网络架构无法满足AI代理的毫秒级发现、即时凭证撤销和加密行为验证等需求。研究分析了三种解决方案：升级现有系统、构建专用注册架构或采用混合模式。由于AI代理的特性要求属于质变而非量变，类似拨号到宽带的转型，研究发现混合方案最具前景——关键代理使用集中式注册，特定场景采用联邦网状架构。尽管升级方案兼容性更好，但全新架构能提供更优性能，只是需要更长的推广时间。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-16T02:33:50Z
- **目录日期**: 2025-06-16
