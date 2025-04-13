# LatentSync: Taming Audio-Conditioned Latent Diffusion Models for Lip Sync with SyncNet Supervision 网页链接本研究提出了一种名为LatentSync的端到端唇同...

**URL**: https://weibo.com/1870858943/Pjot9eGs2

## 原始摘要

LatentSync: Taming Audio-Conditioned Latent Diffusion Models for Lip Sync with SyncNet Supervision <a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.aminer.cn%2Fpub%2F675ba34bae8580e7ff21df01%2Flatentsync-taming-audio-conditioned-latent-diffusion-models-for-lip-sync-with-syncnet" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>本研究提出了一种名为LatentSync的端到端唇同步框架，基于音频条件的潜在扩散模型，无需中间运动表示，与先前的基于像素空间扩散或两阶段生成的唇同步方法不同。该框架能够利用Stable Diffusion的强大能力，直接建模复杂的音频视觉相关性。研究中发现，基于扩散的唇同步方法在不同帧的扩散过程中存在时间一致性不足的问题。为此，研究提出了时间表示对齐（TREPA）方法以增强时间一致性，同时保持唇同步的准确性。TREPA利用大规模自监督视频模型提取的时间表示来对齐生成的帧与真实帧。此外，研究还观察到了SyncNet收敛问题，并进行了全面的实证研究，识别了影响SyncNet收敛的关键因素，包括模型架构、训练超参数和数据预处理方法。通过改变SyncNet的整体训练框架，显著提高了SyncNet的准确性。这些经验也可应用于其他使用SyncNet的唇同步和音频驱动的人像动画方法。基于这些创新，该方法在HDTF和VoxCeleb2数据集上的多个指标上超越了现有的唇同步方法。<a href="https://m.weibo.cn/p/index?extparam=%E4%BA%BA%E5%B7%A5%E6%99%BA%E8%83%BD&amp;containerid=100808f068f0dad74789bee210163c40a4b50d" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">人工智能</span></a><a href="https://m.weibo.cn/p/index?extparam=%E5%A4%A7%E6%A8%A1%E5%9E%8B&amp;containerid=1008082dc9b4e036056e2a00e5499db67ddd30" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">大模型</span></a><a href="https://m.weibo.cn/p/index?extparam=%E7%A1%95%E5%A3%AB%E8%AE%BA%E6%96%87&amp;containerid=1008084cacf38f5903dc7b04550404d0bd3608" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">硕士论文</span></a><img style="" src="https://tvax2.sinaimg.cn/large/6f830abfly1hzn8lwp5a6j21n90z1b29.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

本研究提出LatentSync框架，通过音频条件潜在扩散模型实现端到端唇同步，无需中间运动表示。针对扩散模型的时间一致性问题，创新性地提出TREPA方法，利用自监督视频模型的时间表示增强生成帧的连贯性。同时，研究深入分析了SyncNet的收敛问题，优化了模型架构、训练参数等关键因素，显著提升了SyncNet精度。实验表明，该方法在HDTF和VoxCeleb2数据集上超越现有技术。该成果不仅改进了唇同步效果，其SyncNet优化经验也可推广至其他音频驱动动画任务。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-13T11:05:26Z
- **目录日期**: 2025-04-13
