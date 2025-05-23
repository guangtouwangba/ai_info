# #传统技术让国产视觉基础模型直接上大分##传统技术提效国产AI# 咱就是说啊，视觉基础模型这块儿，国产AI真就是上了个大分——Glint-MVT，来自格灵深瞳的最新成果...

**URL**: https://weibo.com/6105753431/Pt9gfgUF2

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%BC%A0%E7%BB%9F%E6%8A%80%E6%9C%AF%E8%AE%A9%E5%9B%BD%E4%BA%A7%E8%A7%86%E8%A7%89%E5%9F%BA%E7%A1%80%E6%A8%A1%E5%9E%8B%E7%9B%B4%E6%8E%A5%E4%B8%8A%E5%A4%A7%E5%88%86%23&amp;extparam=%23%E4%BC%A0%E7%BB%9F%E6%8A%80%E6%9C%AF%E8%AE%A9%E5%9B%BD%E4%BA%A7%E8%A7%86%E8%A7%89%E5%9F%BA%E7%A1%80%E6%A8%A1%E5%9E%8B%E7%9B%B4%E6%8E%A5%E4%B8%8A%E5%A4%A7%E5%88%86%23" data-hide=""><span class="surl-text">#传统技术让国产视觉基础模型直接上大分#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%BC%A0%E7%BB%9F%E6%8A%80%E6%9C%AF%E6%8F%90%E6%95%88%E5%9B%BD%E4%BA%A7AI%23&amp;extparam=%23%E4%BC%A0%E7%BB%9F%E6%8A%80%E6%9C%AF%E6%8F%90%E6%95%88%E5%9B%BD%E4%BA%A7AI%23" data-hide=""><span class="surl-text">#传统技术提效国产AI#</span></a> <br><br>咱就是说啊，视觉基础模型这块儿，国产AI真就是上了个大分——<br><br>Glint-MVT，来自格灵深瞳的最新成果。<br><br>先来看下成绩——线性探测（LinearProbing）：【图1】。<br><br>简单来说，线性探测是一种测试预训练模型效果的小技巧，测的就是基本功扎不扎实。它的做法是：<br><br>把模型最后一部分换成简单的线性层，其他部分全部保持原样不动；然后只训练这个新加的线性层，通过它的表现来判断模型之前学到的特征好不好用。<br><br>这个测试是在26个分类测试集中跟CLIP和OpenCLIP做了对比，结果显示，国产视觉基础模型平均准确率比OpenCLIP高2.3%，比CLIP高1.1%。<br><br>再来看应用效果。<br><br>如果说视觉基础模型是一个底座，那么它的下游任务，像“图像理解+分割一切”，便是更为直观的效果展现。<br><br>例如下面这张图片，然后我们可以问一下AI：你能提供一个分割掩膜给这个图像中触摸篮球的人吗？【图2】<br><br>很显然，这个任务的难点在于拿篮球的人被其他人的手、身体等部位挡住，分割难度也大幅增加。<br><br>然而，国产AI是不在怕的，啪的一下，超精细地把要求的人物给抠了出来：【图3】。<br><br>我们再来看下更加复杂的案例：【图4】。<br><br>面对如此繁杂的图片，不论要求AI分割哪种食物，它都能精准无误地识别出来。<br><br>Glint-MVT中的MVT，全称叫做Margin-based pretrained Vision Transformer，是团队自研、设计的视觉预训练Transformer模型。<br><br>它的一大亮点，就是创新性地把原先用于人脸识别的间隔Softmax（Margin Softmax）损失函数引入了进来，再通过构造百万级虚拟类别训练模型，显著降低数据噪声影响，提升泛化能力。<br><br>并且从实测和性能效果上来看，在各种专业下游任务中的表现，要比CLIP等其他ViT模型的结果更好。<br><br>在Glint-MVT这个底座之上，团队针对引用表达分割（RES，Referring Expression Segmentation）和图像理解，还分别训练出了多模态模型：Glint-RefSeg模型和MVT-VLM模型。<br><br>Glint-RefSeg是无需特别搜集训练就可以分割一切，从刚才的例子中也是感受到其实力，并且相比其他RES任务的模型，它取得了当前的SOTA。<br><br>至于MVT-VLM，它在图像理解的实力，可以用下面的例子来展示：请分别告诉我图中运动员的性别以及衣服颜色和号码。<br><br>即使图片中红衣服的号码呈现出刁钻的角度，它也能轻松识别：【图5】。<br><br>那么Glint-MVT还有哪些效果，我们继续来看。<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FCXEGGF9tJUycreIpPgV98Q" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">巧妙！一个传统技术让国产视觉基础模型直接上大分</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3ly1i1pemz8z3kj30u00bmmz6.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3ly1i1pen36j5oj30ou0v2hd2.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3ly1i1pencnfuaj30ti14m4i5.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1pennfa0jg30r10lbnpe.gif" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3ly1i1peo38wt7j30u00lztl9.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

格灵深瞳推出的Glint-MVT视觉基础模型在性能测试中表现优异：线性探测准确率分别超过OpenCLIP 2.3%和CLIP 1.1%。该模型创新性地采用人脸识别中的间隔Softmax损失函数，通过百万级虚拟类别训练提升泛化能力。其下游应用Glint-RefSeg在复杂场景分割任务中达到SOTA水平，能精准处理遮挡物体；MVT-VLM多模态模型则展现出强大的图像理解能力，可准确识别细节信息。实验显示，该国产模型在多项视觉任务中的表现优于主流ViT模型，展现了技术创新带来的性能提升。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-23T07:03:09Z
- **目录日期**: 2025-05-23
